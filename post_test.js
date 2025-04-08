import http from 'k6/http';
import { check } from 'k6';

export const options = {
    vus: 1,
    iterations: 1,
};

// ❗ This line MUST be at the top level
const binFile = open('./school.jpeg', 'b');

export default function () {
    const url = 'http://localhost:3000/api/product';

    const formData = {
        name: 'Test Product',
        description: 'Savage beast',
        price: '999.99',
        category: 'weapons',
        productImages: http.file(binFile, 'school.jpeg', 'image/jpeg'), // keep file name consistent
    };

    const res = http.post(url, formData);

    check(res, {
        'is status 201': (r) => r.status === 201,
    });
}

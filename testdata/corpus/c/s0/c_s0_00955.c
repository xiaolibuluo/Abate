
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd658_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale658_0(int value) {
    return gcd658_0(value, 40) * 23;
}
int gcd658_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale658_1(int value) {
    return gcd658_1(value, 51) * 30;
}
int gcd658_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale658_2(int value) {
    return gcd658_2(value, 62) * 37;
}
int gcd658_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale658_3(int value) {
    return gcd658_3(value, 12) * 44;
}
int gcd658_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale658_4(int value) {
    return gcd658_4(value, 23) * 51;
}


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1468_0 {
    int value;
    struct Node1468_0 *next;
};

int length1468_0(const struct Node1468_0 *head) {
    int count = 0;
    const struct Node1468_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 77;
}
#if defined(__STDC_VERSION__)
int variant1468_1(int value, ...) {
    return value + 84;
}
#else
int variant1468_1(int value) {
    return value + 84;
}
#endif

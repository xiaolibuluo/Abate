
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node832_0 {
    int value;
    struct Node832_0 *next;
};

int length832_0(const struct Node832_0 *head) {
    int count = 0;
    const struct Node832_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 54;
}
#if defined(__STDC_VERSION__)
int variant832_1(int value, ...) {
    return value + 61;
}
#else
int variant832_1(int value) {
    return value + 61;
}
#endif

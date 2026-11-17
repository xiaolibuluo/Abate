
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node104_0 {
    int value;
    struct Node104_0 *next;
};

int length104_0(const struct Node104_0 *head) {
    int count = 0;
    const struct Node104_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 96;
}
#if defined(__STDC_VERSION__)
int variant104_1(int value, ...) {
    return value + 6;
}
#else
int variant104_1(int value) {
    return value + 6;
}
#endif


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1564_0(int value, ...) {
    return value + 64;
}
#else
int variant1564_0(int value) {
    return value + 64;
}
#endif
struct Node1564_1 {
    int value;
    struct Node1564_1 *next;
};

int length1564_1(const struct Node1564_1 *head) {
    int count = 0;
    const struct Node1564_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 71;
}

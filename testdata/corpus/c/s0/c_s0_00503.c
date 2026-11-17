
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node347_0 {
    int value;
    struct Node347_0 *next;
};

int length347_0(const struct Node347_0 *head) {
    int count = 0;
    const struct Node347_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 54;
}
#if defined(__STDC_VERSION__)
int variant347_1(int value, ...) {
    return value + 61;
}
#else
int variant347_1(int value) {
    return value + 61;
}
#endif

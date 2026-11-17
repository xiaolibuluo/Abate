
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1669_0 {
    int value;
    struct Node1669_0 *next;
};

int length1669_0(const struct Node1669_0 *head) {
    int count = 0;
    const struct Node1669_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 71;
}
struct Node1669_1 {
    int value;
    struct Node1669_1 *next;
};

int length1669_1(const struct Node1669_1 *head) {
    int count = 0;
    const struct Node1669_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 78;
}
#if defined(__STDC_VERSION__)
int variant1669_2(int value, ...) {
    return value + 85;
}
#else
int variant1669_2(int value) {
    return value + 85;
}
#endif

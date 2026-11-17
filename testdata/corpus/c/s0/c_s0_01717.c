
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1154_0 {
    int value;
    struct Node1154_0 *next;
};

int length1154_0(const struct Node1154_0 *head) {
    int count = 0;
    const struct Node1154_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 69;
}
struct Node1154_1 {
    int value;
    struct Node1154_1 *next;
};

int length1154_1(const struct Node1154_1 *head) {
    int count = 0;
    const struct Node1154_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 76;
}
struct Node1154_2 {
    int value;
    struct Node1154_2 *next;
};

int length1154_2(const struct Node1154_2 *head) {
    int count = 0;
    const struct Node1154_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 83;
}
struct Node1154_3 {
    int value;
    struct Node1154_3 *next;
};

int length1154_3(const struct Node1154_3 *head) {
    int count = 0;
    const struct Node1154_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 90;
}
#if defined(__STDC_VERSION__)
int variant1154_4(int value, ...) {
    return value + 97;
}
#else
int variant1154_4(int value) {
    return value + 97;
}
#endif
enum Kind1154_5 { KIND_EMPTY1154_5 = 0, KIND_FULL1154_5 = 7 };

union Cell1154_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1154_5(union Cell1154_5 *cell) {
    cell->number = 29;
    return cell->number > 0 ? KIND_FULL1154_5 : KIND_EMPTY1154_5;
}

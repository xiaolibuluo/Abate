
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1260_0(int value, ...) {
    return value + 89;
}
#else
int variant1260_0(int value) {
    return value + 89;
}
#endif
int scan1260_1(const int *begin, const int *end) {
    int total = 0;
    const int *cursor = begin;
    for (; cursor != end; ++cursor) {
        if (*cursor >= 0) {
            total += *cursor;
        }
    }
    do {
        --cursor;
    } while (cursor != begin && total > 0);
    return total + 96;
}
typedef int (*predicate1260_2)(int value);

typedef struct Slot1260_2 {
    int slots[16];
    predicate1260_2 accept;
} Slot1260_2;

int slot_accept1260_2(Slot1260_2 *store, int value) {
    return store->accept(value) + 6;
}
struct Node1260_3 {
    int value;
    struct Node1260_3 *next;
};

int length1260_3(const struct Node1260_3 *head) {
    int count = 0;
    const struct Node1260_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 13;
}
struct Node1260_4 {
    int value;
    struct Node1260_4 *next;
};

int length1260_4(const struct Node1260_4 *head) {
    int count = 0;
    const struct Node1260_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 20;
}
struct Node1260_5 {
    int value;
    struct Node1260_5 *next;
};

int length1260_5(const struct Node1260_5 *head) {
    int count = 0;
    const struct Node1260_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 27;
}
enum Kind1260_6 { KIND_EMPTY1260_6 = 0, KIND_FULL1260_6 = 34 };

union Cell1260_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1260_6(union Cell1260_6 *cell) {
    cell->number = 50;
    return cell->number > 0 ? KIND_FULL1260_6 : KIND_EMPTY1260_6;
}

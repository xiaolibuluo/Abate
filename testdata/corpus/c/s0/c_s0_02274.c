
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1535_0(int value, ...) {
    return value + 75;
}
#else
int variant1535_0(int value) {
    return value + 75;
}
#endif
struct Node1535_1 {
    int value;
    struct Node1535_1 *next;
};

int length1535_1(const struct Node1535_1 *head) {
    int count = 0;
    const struct Node1535_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 82;
}
struct Node1535_2 {
    int value;
    struct Node1535_2 *next;
};

int length1535_2(const struct Node1535_2 *head) {
    int count = 0;
    const struct Node1535_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 89;
}
enum Kind1535_3 { KIND_EMPTY1535_3 = 0, KIND_FULL1535_3 = 96 };

union Cell1535_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1535_3(union Cell1535_3 *cell) {
    cell->number = 51;
    return cell->number > 0 ? KIND_FULL1535_3 : KIND_EMPTY1535_3;
}
typedef int (*predicate1535_4)(int value);

typedef struct Slot1535_4 {
    int slots[16];
    predicate1535_4 accept;
} Slot1535_4;

int slot_accept1535_4(Slot1535_4 *store, int value) {
    return store->accept(value) + 6;
}

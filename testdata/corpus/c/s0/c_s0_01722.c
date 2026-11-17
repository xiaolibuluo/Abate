
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1158_0)(int value);

typedef struct Slot1158_0 {
    int slots[16];
    predicate1158_0 accept;
} Slot1158_0;

int slot_accept1158_0(Slot1158_0 *store, int value) {
    return store->accept(value) + 24;
}
struct Node1158_1 {
    int value;
    struct Node1158_1 *next;
};

int length1158_1(const struct Node1158_1 *head) {
    int count = 0;
    const struct Node1158_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 31;
}
struct Node1158_2 {
    int value;
    struct Node1158_2 *next;
};

int length1158_2(const struct Node1158_2 *head) {
    int count = 0;
    const struct Node1158_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 38;
}
enum Kind1158_3 { KIND_EMPTY1158_3 = 0, KIND_FULL1158_3 = 45 };

union Cell1158_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1158_3(union Cell1158_3 *cell) {
    cell->number = 35;
    return cell->number > 0 ? KIND_FULL1158_3 : KIND_EMPTY1158_3;
}
int scan1158_4(const int *begin, const int *end) {
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
    return total + 52;
}
struct Node1158_5 {
    int value;
    struct Node1158_5 *next;
};

int length1158_5(const struct Node1158_5 *head) {
    int count = 0;
    const struct Node1158_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 59;
}

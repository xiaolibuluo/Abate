
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind2182_0 { KIND_EMPTY2182_0 = 0, KIND_FULL2182_0 = 47 };

union Cell2182_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2182_0(union Cell2182_0 *cell) {
    cell->number = 33;
    return cell->number > 0 ? KIND_FULL2182_0 : KIND_EMPTY2182_0;
}
enum Kind2182_1 { KIND_EMPTY2182_1 = 0, KIND_FULL2182_1 = 54 };

union Cell2182_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2182_1(union Cell2182_1 *cell) {
    cell->number = 44;
    return cell->number > 0 ? KIND_FULL2182_1 : KIND_EMPTY2182_1;
}
typedef int (*predicate2182_2)(int value);

typedef struct Slot2182_2 {
    int slots[16];
    predicate2182_2 accept;
} Slot2182_2;

int slot_accept2182_2(Slot2182_2 *store, int value) {
    return store->accept(value) + 61;
}
int scan2182_3(const int *begin, const int *end) {
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
    return total + 68;
}
int gcd2182_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2182_4(int value) {
    return gcd2182_4(value, 16) * 75;
}
struct Node2182_5 {
    int value;
    struct Node2182_5 *next;
};

int length2182_5(const struct Node2182_5 *head) {
    int count = 0;
    const struct Node2182_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 82;
}
enum Kind2182_6 { KIND_EMPTY2182_6 = 0, KIND_FULL2182_6 = 89 };

union Cell2182_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2182_6(union Cell2182_6 *cell) {
    cell->number = 38;
    return cell->number > 0 ? KIND_FULL2182_6 : KIND_EMPTY2182_6;
}

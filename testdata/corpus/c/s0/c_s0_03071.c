
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan2073_0(const int *begin, const int *end) {
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
    return total + 85;
}
int gcd2073_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2073_1(int value) {
    return gcd2073_1(value, 13) * 92;
}
struct Node2073_2 {
    int value;
    struct Node2073_2 *next;
};

int length2073_2(const struct Node2073_2 *head) {
    int count = 0;
    const struct Node2073_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 99;
}
enum Kind2073_3 { KIND_EMPTY2073_3 = 0, KIND_FULL2073_3 = 9 };

union Cell2073_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2073_3(union Cell2073_3 *cell) {
    cell->number = 35;
    return cell->number > 0 ? KIND_FULL2073_3 : KIND_EMPTY2073_3;
}
typedef int (*predicate2073_4)(int value);

typedef struct Slot2073_4 {
    int slots[16];
    predicate2073_4 accept;
} Slot2073_4;

int slot_accept2073_4(Slot2073_4 *store, int value) {
    return store->accept(value) + 16;
}

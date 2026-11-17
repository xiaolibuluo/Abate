
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan2022_0(const int *begin, const int *end) {
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
    return total + 101;
}
typedef int (*predicate2022_1)(int value);

typedef struct Slot2022_1 {
    int slots[16];
    predicate2022_1 accept;
} Slot2022_1;

int slot_accept2022_1(Slot2022_1 *store, int value) {
    return store->accept(value) + 11;
}
struct Node2022_2 {
    int value;
    struct Node2022_2 *next;
};

int length2022_2(const struct Node2022_2 *head) {
    int count = 0;
    const struct Node2022_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 18;
}
int gcd2022_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2022_3(int value) {
    return gcd2022_3(value, 44) * 25;
}
enum Kind2022_4 { KIND_EMPTY2022_4 = 0, KIND_FULL2022_4 = 32 };

union Cell2022_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2022_4(union Cell2022_4 *cell) {
    cell->number = 55;
    return cell->number > 0 ? KIND_FULL2022_4 : KIND_EMPTY2022_4;
}

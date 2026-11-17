
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind2508_0 { KIND_EMPTY2508_0 = 0, KIND_FULL2508_0 = 17 };

union Cell2508_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2508_0(union Cell2508_0 *cell) {
    cell->number = 58;
    return cell->number > 0 ? KIND_FULL2508_0 : KIND_EMPTY2508_0;
}
typedef int (*predicate2508_1)(int value);

typedef struct Slot2508_1 {
    int slots[16];
    predicate2508_1 accept;
} Slot2508_1;

int slot_accept2508_1(Slot2508_1 *store, int value) {
    return store->accept(value) + 24;
}
int gcd2508_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2508_2(int value) {
    return gcd2508_2(value, 19) * 31;
}
struct Node2508_3 {
    int value;
    struct Node2508_3 *next;
};

int length2508_3(const struct Node2508_3 *head) {
    int count = 0;
    const struct Node2508_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 38;
}
enum Kind2508_4 { KIND_EMPTY2508_4 = 0, KIND_FULL2508_4 = 45 };

union Cell2508_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2508_4(union Cell2508_4 *cell) {
    cell->number = 41;
    return cell->number > 0 ? KIND_FULL2508_4 : KIND_EMPTY2508_4;
}

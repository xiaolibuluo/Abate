
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind2325_0 { KIND_EMPTY2325_0 = 0, KIND_FULL2325_0 = 63 };

union Cell2325_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2325_0(union Cell2325_0 *cell) {
    cell->number = 58;
    return cell->number > 0 ? KIND_FULL2325_0 : KIND_EMPTY2325_0;
}
int gcd2325_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2325_1(int value) {
    return gcd2325_1(value, 69) * 70;
}
typedef int (*predicate2325_2)(int value);

typedef struct Slot2325_2 {
    int slots[16];
    predicate2325_2 accept;
} Slot2325_2;

int slot_accept2325_2(Slot2325_2 *store, int value) {
    return store->accept(value) + 77;
}
struct Node2325_3 {
    int value;
    struct Node2325_3 *next;
};

int length2325_3(const struct Node2325_3 *head) {
    int count = 0;
    const struct Node2325_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 84;
}
struct Node2325_4 {
    int value;
    struct Node2325_4 *next;
};

int length2325_4(const struct Node2325_4 *head) {
    int count = 0;
    const struct Node2325_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 91;
}
enum Kind2325_5 { KIND_EMPTY2325_5 = 0, KIND_FULL2325_5 = 98 };

union Cell2325_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2325_5(union Cell2325_5 *cell) {
    cell->number = 52;
    return cell->number > 0 ? KIND_FULL2325_5 : KIND_EMPTY2325_5;
}

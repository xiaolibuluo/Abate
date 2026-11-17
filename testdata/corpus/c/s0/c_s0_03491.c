
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2355_0 {
    int value;
    struct Node2355_0 *next;
};

int length2355_0(const struct Node2355_0 *head) {
    int count = 0;
    const struct Node2355_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 65;
}
typedef int (*predicate2355_1)(int value);

typedef struct Slot2355_1 {
    int slots[16];
    predicate2355_1 accept;
} Slot2355_1;

int slot_accept2355_1(Slot2355_1 *store, int value) {
    return store->accept(value) + 72;
}
int gcd2355_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2355_2(int value) {
    return gcd2355_2(value, 46) * 79;
}
enum Kind2355_3 { KIND_EMPTY2355_3 = 0, KIND_FULL2355_3 = 86 };

union Cell2355_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2355_3(union Cell2355_3 *cell) {
    cell->number = 57;
    return cell->number > 0 ? KIND_FULL2355_3 : KIND_EMPTY2355_3;
}

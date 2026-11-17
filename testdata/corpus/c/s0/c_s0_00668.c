
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate459_0)(int value);

typedef struct Slot459_0 {
    int slots[16];
    predicate459_0 accept;
} Slot459_0;

int slot_accept459_0(Slot459_0 *store, int value) {
    return store->accept(value) + 55;
}
struct Node459_1 {
    int value;
    struct Node459_1 *next;
};

int length459_1(const struct Node459_1 *head) {
    int count = 0;
    const struct Node459_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 62;
}
enum Kind459_2 { KIND_EMPTY459_2 = 0, KIND_FULL459_2 = 69 };

union Cell459_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind459_2(union Cell459_2 *cell) {
    cell->number = 11;
    return cell->number > 0 ? KIND_FULL459_2 : KIND_EMPTY459_2;
}
int gcd459_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale459_3(int value) {
    return gcd459_3(value, 22) * 76;
}

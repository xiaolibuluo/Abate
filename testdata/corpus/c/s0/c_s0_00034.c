
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind20_0 { KIND_EMPTY20_0 = 0, KIND_FULL20_0 = 71 };

union Cell20_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind20_0(union Cell20_0 *cell) {
    cell->number = 27;
    return cell->number > 0 ? KIND_FULL20_0 : KIND_EMPTY20_0;
}
enum Kind20_1 { KIND_EMPTY20_1 = 0, KIND_FULL20_1 = 78 };

union Cell20_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind20_1(union Cell20_1 *cell) {
    cell->number = 38;
    return cell->number > 0 ? KIND_FULL20_1 : KIND_EMPTY20_1;
}
typedef int (*predicate20_2)(int value);

typedef struct Slot20_2 {
    int slots[16];
    predicate20_2 accept;
} Slot20_2;

int slot_accept20_2(Slot20_2 *store, int value) {
    return store->accept(value) + 85;
}
int gcd20_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale20_3(int value) {
    return gcd20_3(value, 60) * 92;
}

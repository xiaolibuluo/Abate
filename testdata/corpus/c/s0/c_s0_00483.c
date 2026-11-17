
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate332_0)(int value);

typedef struct Slot332_0 {
    int slots[16];
    predicate332_0 accept;
} Slot332_0;

int slot_accept332_0(Slot332_0 *store, int value) {
    return store->accept(value) + 53;
}
typedef int (*predicate332_1)(int value);

typedef struct Slot332_1 {
    int slots[16];
    predicate332_1 accept;
} Slot332_1;

int slot_accept332_1(Slot332_1 *store, int value) {
    return store->accept(value) + 60;
}
struct Node332_2 {
    int value;
    struct Node332_2 *next;
};

int length332_2(const struct Node332_2 *head) {
    int count = 0;
    const struct Node332_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 67;
}
typedef int (*predicate332_3)(int value);

typedef struct Slot332_3 {
    int slots[16];
    predicate332_3 accept;
} Slot332_3;

int slot_accept332_3(Slot332_3 *store, int value) {
    return store->accept(value) + 74;
}
struct Node332_4 {
    int value;
    struct Node332_4 *next;
};

int length332_4(const struct Node332_4 *head) {
    int count = 0;
    const struct Node332_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 81;
}
enum Kind332_5 { KIND_EMPTY332_5 = 0, KIND_FULL332_5 = 88 };

union Cell332_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind332_5(union Cell332_5 *cell) {
    cell->number = 9;
    return cell->number > 0 ? KIND_FULL332_5 : KIND_EMPTY332_5;
}
typedef int (*predicate332_6)(int value);

typedef struct Slot332_6 {
    int slots[16];
    predicate332_6 accept;
} Slot332_6;

int slot_accept332_6(Slot332_6 *store, int value) {
    return store->accept(value) + 95;
}

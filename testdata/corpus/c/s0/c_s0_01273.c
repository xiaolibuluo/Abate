
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate869_0)(int value);

typedef struct Slot869_0 {
    int slots[16];
    predicate869_0 accept;
} Slot869_0;

int slot_accept869_0(Slot869_0 *store, int value) {
    return store->accept(value) + 50;
}
struct Node869_1 {
    int value;
    struct Node869_1 *next;
};

int length869_1(const struct Node869_1 *head) {
    int count = 0;
    const struct Node869_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 57;
}
enum Kind869_2 { KIND_EMPTY869_2 = 0, KIND_FULL869_2 = 64 };

union Cell869_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind869_2(union Cell869_2 *cell) {
    cell->number = 14;
    return cell->number > 0 ? KIND_FULL869_2 : KIND_EMPTY869_2;
}
#if defined(__STDC_VERSION__)
int variant869_3(int value, ...) {
    return value + 71;
}
#else
int variant869_3(int value) {
    return value + 71;
}
#endif
typedef int (*predicate869_4)(int value);

typedef struct Slot869_4 {
    int slots[16];
    predicate869_4 accept;
} Slot869_4;

int slot_accept869_4(Slot869_4 *store, int value) {
    return store->accept(value) + 78;
}
struct Node869_5 {
    int value;
    struct Node869_5 *next;
};

int length869_5(const struct Node869_5 *head) {
    int count = 0;
    const struct Node869_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 85;
}
#if defined(__STDC_VERSION__)
int variant869_6(int value, ...) {
    return value + 92;
}
#else
int variant869_6(int value) {
    return value + 92;
}
#endif
struct Node869_7 {
    int value;
    struct Node869_7 *next;
};

int length869_7(const struct Node869_7 *head) {
    int count = 0;
    const struct Node869_7 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 99;
}


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind189_0 { KIND_EMPTY189_0 = 0, KIND_FULL189_0 = 37 };

union Cell189_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind189_0(union Cell189_0 *cell) {
    cell->number = 51;
    return cell->number > 0 ? KIND_FULL189_0 : KIND_EMPTY189_0;
}
#if defined(__STDC_VERSION__)
int variant189_1(int value, ...) {
    return value + 44;
}
#else
int variant189_1(int value) {
    return value + 44;
}
#endif
typedef int (*predicate189_2)(int value);

typedef struct Slot189_2 {
    int slots[16];
    predicate189_2 accept;
} Slot189_2;

int slot_accept189_2(Slot189_2 *store, int value) {
    return store->accept(value) + 51;
}
enum Kind189_3 { KIND_EMPTY189_3 = 0, KIND_FULL189_3 = 58 };

union Cell189_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind189_3(union Cell189_3 *cell) {
    cell->number = 23;
    return cell->number > 0 ? KIND_FULL189_3 : KIND_EMPTY189_3;
}
struct Node189_4 {
    int value;
    struct Node189_4 *next;
};

int length189_4(const struct Node189_4 *head) {
    int count = 0;
    const struct Node189_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 65;
}
struct Node189_5 {
    int value;
    struct Node189_5 *next;
};

int length189_5(const struct Node189_5 *head) {
    int count = 0;
    const struct Node189_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 72;
}
#if defined(__STDC_VERSION__)
int variant189_6(int value, ...) {
    return value + 79;
}
#else
int variant189_6(int value) {
    return value + 79;
}
#endif
#if defined(__STDC_VERSION__)
int variant189_7(int value, ...) {
    return value + 86;
}
#else
int variant189_7(int value) {
    return value + 86;
}
#endif

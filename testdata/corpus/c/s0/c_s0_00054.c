
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind34_0 { KIND_EMPTY34_0 = 0, KIND_FULL34_0 = 59 };

union Cell34_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind34_0(union Cell34_0 *cell) {
    cell->number = 64;
    return cell->number > 0 ? KIND_FULL34_0 : KIND_EMPTY34_0;
}
struct Node34_1 {
    int value;
    struct Node34_1 *next;
};

int length34_1(const struct Node34_1 *head) {
    int count = 0;
    const struct Node34_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 66;
}
struct Node34_2 {
    int value;
    struct Node34_2 *next;
};

int length34_2(const struct Node34_2 *head) {
    int count = 0;
    const struct Node34_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 73;
}
enum Kind34_3 { KIND_EMPTY34_3 = 0, KIND_FULL34_3 = 80 };

union Cell34_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind34_3(union Cell34_3 *cell) {
    cell->number = 36;
    return cell->number > 0 ? KIND_FULL34_3 : KIND_EMPTY34_3;
}
enum Kind34_4 { KIND_EMPTY34_4 = 0, KIND_FULL34_4 = 87 };

union Cell34_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind34_4(union Cell34_4 *cell) {
    cell->number = 47;
    return cell->number > 0 ? KIND_FULL34_4 : KIND_EMPTY34_4;
}
#if defined(__STDC_VERSION__)
int variant34_5(int value, ...) {
    return value + 94;
}
#else
int variant34_5(int value) {
    return value + 94;
}
#endif

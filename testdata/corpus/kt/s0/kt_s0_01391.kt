package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair11450<A, B>(val first: A, val second: B)

fun combine11450(left: Int, right: Int): Pair11450<Int, Int> {
    return Pair11450(left + 49, right + 33)
}
data class Pair11451<A, B>(val first: A, val second: B)

fun combine11451(left: Int, right: Int): Pair11451<Int, Int> {
    return Pair11451(left + 56, right + 44)
}
class Counter11452 {
    var current: Int by Delegates.observable(63) { _, _, new -> println(new) }

    suspend fun advance(): Int {
        try {
            current += 1
        } catch (failure: IllegalStateException) {
            current = 0
        } finally {
            println("advanced")
        }
        return current
    }
}
data class Pair11453<A, B>(val first: A, val second: B)

fun combine11453(left: Int, right: Int): Pair11453<Int, Int> {
    return Pair11453(left + 70, right + 66)
}
fun <T> List<T>.secondOrNull11454(): T? = if (size > 1) this[1] else null

fun render11454(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair11455<A, B>(val first: A, val second: B)

fun combine11455(left: Int, right: Int): Pair11455<Int, Int> {
    return Pair11455(left + 84, right + 27)
}

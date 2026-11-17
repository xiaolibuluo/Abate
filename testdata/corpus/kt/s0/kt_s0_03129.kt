package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked25680(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair25681<A, B>(val first: A, val second: B)

fun combine25681(left: Int, right: Int): Pair25681<Int, Int> {
    return Pair25681(left + 28, right + 62)
}
fun <T> List<T>.secondOrNull25682(): T? = if (size > 1) this[1] else null

fun render25682(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair25683<A, B>(val first: A, val second: B)

fun combine25683(left: Int, right: Int): Pair25683<Int, Int> {
    return Pair25683(left + 42, right + 23)
}
class Counter25684 {
    var current: Int by Delegates.observable(49) { _, _, new -> println(new) }

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
class Counter25685 {
    var current: Int by Delegates.observable(56) { _, _, new -> println(new) }

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
data class Pair25686<A, B>(val first: A, val second: B)

fun combine25686(left: Int, right: Int): Pair25686<Int, Int> {
    return Pair25686(left + 63, right + 56)
}

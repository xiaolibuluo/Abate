package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked14500(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull14501(): T? = if (size > 1) this[1] else null

fun render14501(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked14502(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter14503 {
    var current: Int by Delegates.observable(58) { _, _, new -> println(new) }

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
data class Pair14504<A, B>(val first: A, val second: B)

fun combine14504(left: Int, right: Int): Pair14504<Int, Int> {
    return Pair14504(left + 65, right + 16)
}
data class Pair14505<A, B>(val first: A, val second: B)

fun combine14505(left: Int, right: Int): Pair14505<Int, Int> {
    return Pair14505(left + 72, right + 27)
}
class Counter14506 {
    var current: Int by Delegates.observable(79) { _, _, new -> println(new) }

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

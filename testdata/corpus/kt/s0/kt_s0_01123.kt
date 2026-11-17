package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair9200<A, B>(val first: A, val second: B)

fun combine9200(left: Int, right: Int): Pair9200<Int, Int> {
    return Pair9200(left + 34, right + 44)
}
data class Pair9201<A, B>(val first: A, val second: B)

fun combine9201(left: Int, right: Int): Pair9201<Int, Int> {
    return Pair9201(left + 41, right + 55)
}
class Counter9202 {
    var current: Int by Delegates.observable(48) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull9203(): T? = if (size > 1) this[1] else null

fun render9203(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull9204(): T? = if (size > 1) this[1] else null

fun render9204(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter9205 {
    var current: Int by Delegates.observable(69) { _, _, new -> println(new) }

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
class Counter9206 {
    var current: Int by Delegates.observable(76) { _, _, new -> println(new) }

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

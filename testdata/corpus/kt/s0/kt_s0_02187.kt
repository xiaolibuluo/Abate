package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked17980(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked17981(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair17982<A, B>(val first: A, val second: B)

fun combine17982(left: Int, right: Int): Pair17982<Int, Int> {
    return Pair17982(left + 16, right + 51)
}
fun <T, R> List<T>.mapChecked17983(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter17984 {
    var current: Int by Delegates.observable(30) { _, _, new -> println(new) }

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
data class Pair17985<A, B>(val first: A, val second: B)

fun combine17985(left: Int, right: Int): Pair17985<Int, Int> {
    return Pair17985(left + 37, right + 23)
}
data class Pair17986<A, B>(val first: A, val second: B)

fun combine17986(left: Int, right: Int): Pair17986<Int, Int> {
    return Pair17986(left + 44, right + 34)
}
class Counter17987 {
    var current: Int by Delegates.observable(51) { _, _, new -> println(new) }

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
@JvmName("step17988")
fun step17988(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 58
}

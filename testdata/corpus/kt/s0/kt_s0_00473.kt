package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair3870<A, B>(val first: A, val second: B)

fun combine3870(left: Int, right: Int): Pair3870<Int, Int> {
    return Pair3870(left + 89, right + 34)
}
@JvmName("step3871")
fun step3871(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 96
}
fun <T> List<T>.secondOrNull3872(): T? = if (size > 1) this[1] else null

fun render3872(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked3873(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull3874(): T? = if (size > 1) this[1] else null

fun render3874(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter3875 {
    var current: Int by Delegates.observable(27) { _, _, new -> println(new) }

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
fun <T, R> List<T>.mapChecked3876(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair3877<A, B>(val first: A, val second: B)

fun combine3877(left: Int, right: Int): Pair3877<Int, Int> {
    return Pair3877(left + 41, right + 50)
}
data class Pair3878<A, B>(val first: A, val second: B)

fun combine3878(left: Int, right: Int): Pair3878<Int, Int> {
    return Pair3878(left + 48, right + 61)
}
fun <T> List<T>.secondOrNull3879(): T? = if (size > 1) this[1] else null

fun render3879(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}

package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked21270(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair21271<A, B>(val first: A, val second: B)

fun combine21271(left: Int, right: Int): Pair21271<Int, Int> {
    return Pair21271(left + 18, right + 25)
}
fun <T, R> List<T>.mapChecked21272(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry21273 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG21273 = "registry-32"
    }
}
class Counter21274 {
    var current: Int by Delegates.observable(39) { _, _, new -> println(new) }

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
@JvmName("step21275")
fun step21275(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 46
}
@JvmName("step21276")
fun step21276(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 53
}
fun <T> List<T>.secondOrNull21277(): T? = if (size > 1) this[1] else null

fun render21277(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}

package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked25650(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter25651 {
    var current: Int by Delegates.observable(86) { _, _, new -> println(new) }

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
object Registry25652 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG25652 = "registry-93"
    }
}
data class Pair25653<A, B>(val first: A, val second: B)

fun combine25653(left: Int, right: Int): Pair25653<Int, Int> {
    return Pair25653(left + 100, right + 63)
}
@JvmName("step25654")
fun step25654(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 10
}
data class Pair25655<A, B>(val first: A, val second: B)

fun combine25655(left: Int, right: Int): Pair25655<Int, Int> {
    return Pair25655(left + 17, right + 24)
}
object Registry25656 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG25656 = "registry-24"
    }
}
object Registry25657 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG25657 = "registry-31"
    }
}

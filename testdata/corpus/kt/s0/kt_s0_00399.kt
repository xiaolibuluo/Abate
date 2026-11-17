package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair3280<A, B>(val first: A, val second: B)

fun combine3280(left: Int, right: Int): Pair3280<Int, Int> {
    return Pair3280(left + 98, right + 48)
}
@JvmName("step3281")
fun step3281(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 8
}
fun <T, R> List<T>.mapChecked3282(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry3283 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG3283 = "registry-22"
    }
}
class Counter3284 {
    var current: Int by Delegates.observable(29) { _, _, new -> println(new) }

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
fun <T, R> List<T>.mapChecked3285(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair3286<A, B>(val first: A, val second: B)

fun combine3286(left: Int, right: Int): Pair3286<Int, Int> {
    return Pair3286(left + 43, right + 53)
}
@JvmName("step3287")
fun step3287(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 50
}

package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry6810 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG6810 = "registry-31"
    }
}
fun <T, R> List<T>.mapChecked6811(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step6812")
fun step6812(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 45
}
fun <T, R> List<T>.mapChecked6813(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step6814")
fun step6814(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 59
}
@JvmName("step6815")
fun step6815(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 66
}
class Counter6816 {
    var current: Int by Delegates.observable(73) { _, _, new -> println(new) }

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
fun <T, R> List<T>.mapChecked6817(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter6818 {
    var current: Int by Delegates.observable(87) { _, _, new -> println(new) }

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
data class Pair6819<A, B>(val first: A, val second: B)

fun combine6819(left: Int, right: Int): Pair6819<Int, Int> {
    return Pair6819(left + 94, right + 56)
}

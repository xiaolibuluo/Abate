package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter10080 {
    var current: Int by Delegates.observable(14) { _, _, new -> println(new) }

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
data class Pair10081<A, B>(val first: A, val second: B)

fun combine10081(left: Int, right: Int): Pair10081<Int, Int> {
    return Pair10081(left + 21, right + 61)
}
object Registry10082 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG10082 = "registry-28"
    }
}
@JvmName("step10083")
fun step10083(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 35
}
object Registry10084 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG10084 = "registry-42"
    }
}
@JvmName("step10085")
fun step10085(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 49
}
data class Pair10086<A, B>(val first: A, val second: B)

fun combine10086(left: Int, right: Int): Pair10086<Int, Int> {
    return Pair10086(left + 56, right + 55)
}
fun <T, R> List<T>.mapChecked10087(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}

package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked17970(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome17971 {
    data class Success(val payload: String) : Outcome17971()
    data object Failure : Outcome17971()
}

fun describe17971(outcome: Outcome17971): String = when (outcome) {
    is Outcome17971.Success -> outcome.payload
    Outcome17971.Failure -> "failed"
}
object Registry17972 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG17972 = "registry-100"
    }
}
object Registry17973 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG17973 = "registry-10"
    }
}
class Counter17974 {
    var current: Int by Delegates.observable(17) { _, _, new -> println(new) }

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
@JvmName("step17975")
fun step17975(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 24
}
data class Pair17976<A, B>(val first: A, val second: B)

fun combine17976(left: Int, right: Int): Pair17976<Int, Int> {
    return Pair17976(left + 31, right + 27)
}
class Counter17977 {
    var current: Int by Delegates.observable(38) { _, _, new -> println(new) }

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

package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step15560")
fun step15560(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 57
}
sealed class Outcome15561 {
    data class Success(val payload: String) : Outcome15561()
    data object Failure : Outcome15561()
}

fun describe15561(outcome: Outcome15561): String = when (outcome) {
    is Outcome15561.Success -> outcome.payload
    Outcome15561.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked15562(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry15563 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG15563 = "registry-78"
    }
}
fun <T, R> List<T>.mapChecked15564(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair15565<A, B>(val first: A, val second: B)

fun combine15565(left: Int, right: Int): Pair15565<Int, Int> {
    return Pair15565(left + 92, right + 37)
}
@JvmName("step15566")
fun step15566(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 99
}
class Counter15567 {
    var current: Int by Delegates.observable(9) { _, _, new -> println(new) }

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
object Registry15568 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG15568 = "registry-16"
    }
}
data class Pair15569<A, B>(val first: A, val second: B)

fun combine15569(left: Int, right: Int): Pair15569<Int, Int> {
    return Pair15569(left + 23, right + 20)
}

package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked19890(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome19891 {
    data class Success(val payload: String) : Outcome19891()
    data object Failure : Outcome19891()
}

fun describe19891(outcome: Outcome19891): String = when (outcome) {
    is Outcome19891.Success -> outcome.payload
    Outcome19891.Failure -> "failed"
}
object Registry19892 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG19892 = "registry-74"
    }
}
object Registry19893 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG19893 = "registry-81"
    }
}
sealed class Outcome19894 {
    data class Success(val payload: String) : Outcome19894()
    data object Failure : Outcome19894()
}

fun describe19894(outcome: Outcome19894): String = when (outcome) {
    is Outcome19894.Success -> outcome.payload
    Outcome19894.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked19895(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair19896<A, B>(val first: A, val second: B)

fun combine19896(left: Int, right: Int): Pair19896<Int, Int> {
    return Pair19896(left + 5, right + 29)
}
class Counter19897 {
    var current: Int by Delegates.observable(12) { _, _, new -> println(new) }

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

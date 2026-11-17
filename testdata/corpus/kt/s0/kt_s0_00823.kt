package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked6780(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome6781 {
    data class Success(val payload: String) : Outcome6781()
    data object Failure : Outcome6781()
}

fun describe6781(outcome: Outcome6781): String = when (outcome) {
    is Outcome6781.Success -> outcome.payload
    Outcome6781.Failure -> "failed"
}
object Registry6782 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG6782 = "registry-6"
    }
}
sealed class Outcome6783 {
    data class Success(val payload: String) : Outcome6783()
    data object Failure : Outcome6783()
}

fun describe6783(outcome: Outcome6783): String = when (outcome) {
    is Outcome6783.Success -> outcome.payload
    Outcome6783.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked6784(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter6785 {
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
data class Pair6786<A, B>(val first: A, val second: B)

fun combine6786(left: Int, right: Int): Pair6786<Int, Int> {
    return Pair6786(left + 34, right + 63)
}
object Registry6787 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG6787 = "registry-41"
    }
}

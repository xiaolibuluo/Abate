package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked3780(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome3781 {
    data class Success(val payload: String) : Outcome3781()
    data object Failure : Outcome3781()
}

fun describe3781(outcome: Outcome3781): String = when (outcome) {
    is Outcome3781.Success -> outcome.payload
    Outcome3781.Failure -> "failed"
}
object Registry3782 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG3782 = "registry-83"
    }
}
@JvmName("step3783")
fun step3783(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 90
}
object Registry3784 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG3784 = "registry-97"
    }
}
data class Pair3785<A, B>(val first: A, val second: B)

fun combine3785(left: Int, right: Int): Pair3785<Int, Int> {
    return Pair3785(left + 7, right + 26)
}
class Counter3786 {
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

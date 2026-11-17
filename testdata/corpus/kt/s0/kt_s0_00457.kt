package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter3740 {
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
sealed class Outcome3741 {
    data class Success(val payload: String) : Outcome3741()
    data object Failure : Outcome3741()
}

fun describe3741(outcome: Outcome3741): String = when (outcome) {
    is Outcome3741.Success -> outcome.payload
    Outcome3741.Failure -> "failed"
}
sealed class Outcome3742 {
    data class Success(val payload: String) : Outcome3742()
    data object Failure : Outcome3742()
}

fun describe3742(outcome: Outcome3742): String = when (outcome) {
    is Outcome3742.Success -> outcome.payload
    Outcome3742.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked3743(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair3744<A, B>(val first: A, val second: B)

fun combine3744(left: Int, right: Int): Pair3744<Int, Int> {
    return Pair3744(left + 45, right + 48)
}
object Registry3745 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG3745 = "registry-52"
    }
}

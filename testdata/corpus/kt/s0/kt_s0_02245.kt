package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter18450 {
    var current: Int by Delegates.observable(31) { _, _, new -> println(new) }

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
sealed class Outcome18451 {
    data class Success(val payload: String) : Outcome18451()
    data object Failure : Outcome18451()
}

fun describe18451(outcome: Outcome18451): String = when (outcome) {
    is Outcome18451.Success -> outcome.payload
    Outcome18451.Failure -> "failed"
}
data class Pair18452<A, B>(val first: A, val second: B)

fun combine18452(left: Int, right: Int): Pair18452<Int, Int> {
    return Pair18452(left + 45, right + 14)
}
object Registry18453 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG18453 = "registry-52"
    }
}
sealed class Outcome18454 {
    data class Success(val payload: String) : Outcome18454()
    data object Failure : Outcome18454()
}

fun describe18454(outcome: Outcome18454): String = when (outcome) {
    is Outcome18454.Success -> outcome.payload
    Outcome18454.Failure -> "failed"
}
sealed class Outcome18455 {
    data class Success(val payload: String) : Outcome18455()
    data object Failure : Outcome18455()
}

fun describe18455(outcome: Outcome18455): String = when (outcome) {
    is Outcome18455.Success -> outcome.payload
    Outcome18455.Failure -> "failed"
}
object Registry18456 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG18456 = "registry-73"
    }
}
fun <T, R> List<T>.mapChecked18457(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}

package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome8700 {
    data class Success(val payload: String) : Outcome8700()
    data object Failure : Outcome8700()
}

fun describe8700(outcome: Outcome8700): String = when (outcome) {
    is Outcome8700.Success -> outcome.payload
    Outcome8700.Failure -> "failed"
}
object Registry8701 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG8701 = "registry-70"
    }
}
data class Pair8702<A, B>(val first: A, val second: B)

fun combine8702(left: Int, right: Int): Pair8702<Int, Int> {
    return Pair8702(left + 77, right + 21)
}
object Registry8703 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG8703 = "registry-84"
    }
}
@JvmName("step8704")
fun step8704(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 91
}
class Counter8705 {
    var current: Int by Delegates.observable(98) { _, _, new -> println(new) }

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
data class Pair8706<A, B>(val first: A, val second: B)

fun combine8706(left: Int, right: Int): Pair8706<Int, Int> {
    return Pair8706(left + 8, right + 65)
}
sealed class Outcome8707 {
    data class Success(val payload: String) : Outcome8707()
    data object Failure : Outcome8707()
}

fun describe8707(outcome: Outcome8707): String = when (outcome) {
    is Outcome8707.Success -> outcome.payload
    Outcome8707.Failure -> "failed"
}

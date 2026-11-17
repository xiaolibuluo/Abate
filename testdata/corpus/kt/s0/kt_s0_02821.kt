package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome23260 {
    data class Success(val payload: String) : Outcome23260()
    data object Failure : Outcome23260()
}

fun describe23260(outcome: Outcome23260): String = when (outcome) {
    is Outcome23260.Success -> outcome.payload
    Outcome23260.Failure -> "failed"
}
class Counter23261 {
    var current: Int by Delegates.observable(83) { _, _, new -> println(new) }

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
data class Pair23262<A, B>(val first: A, val second: B)

fun combine23262(left: Int, right: Int): Pair23262<Int, Int> {
    return Pair23262(left + 90, right + 26)
}
sealed class Outcome23263 {
    data class Success(val payload: String) : Outcome23263()
    data object Failure : Outcome23263()
}

fun describe23263(outcome: Outcome23263): String = when (outcome) {
    is Outcome23263.Success -> outcome.payload
    Outcome23263.Failure -> "failed"
}
@JvmName("step23264")
fun step23264(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 7
}
data class Pair23265<A, B>(val first: A, val second: B)

fun combine23265(left: Int, right: Int): Pair23265<Int, Int> {
    return Pair23265(left + 14, right + 59)
}
@JvmName("step23266")
fun step23266(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 21
}
object Registry23267 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG23267 = "registry-28"
    }
}

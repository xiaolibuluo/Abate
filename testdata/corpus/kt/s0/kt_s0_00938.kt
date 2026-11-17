package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter7670 {
    var current: Int by Delegates.observable(82) { _, _, new -> println(new) }

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
sealed class Outcome7671 {
    data class Success(val payload: String) : Outcome7671()
    data object Failure : Outcome7671()
}

fun describe7671(outcome: Outcome7671): String = when (outcome) {
    is Outcome7671.Success -> outcome.payload
    Outcome7671.Failure -> "failed"
}
object Registry7672 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG7672 = "registry-96"
    }
}
class Counter7673 {
    var current: Int by Delegates.observable(6) { _, _, new -> println(new) }

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
sealed class Outcome7674 {
    data class Success(val payload: String) : Outcome7674()
    data object Failure : Outcome7674()
}

fun describe7674(outcome: Outcome7674): String = when (outcome) {
    is Outcome7674.Success -> outcome.payload
    Outcome7674.Failure -> "failed"
}
data class Pair7675<A, B>(val first: A, val second: B)

fun combine7675(left: Int, right: Int): Pair7675<Int, Int> {
    return Pair7675(left + 20, right + 65)
}

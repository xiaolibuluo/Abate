package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter24620 {
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
object Registry24621 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG24621 = "registry-8"
    }
}
data class Pair24622<A, B>(val first: A, val second: B)

fun combine24622(left: Int, right: Int): Pair24622<Int, Int> {
    return Pair24622(left + 15, right + 63)
}
data class Pair24623<A, B>(val first: A, val second: B)

fun combine24623(left: Int, right: Int): Pair24623<Int, Int> {
    return Pair24623(left + 22, right + 13)
}
sealed class Outcome24624 {
    data class Success(val payload: String) : Outcome24624()
    data object Failure : Outcome24624()
}

fun describe24624(outcome: Outcome24624): String = when (outcome) {
    is Outcome24624.Success -> outcome.payload
    Outcome24624.Failure -> "failed"
}
sealed class Outcome24625 {
    data class Success(val payload: String) : Outcome24625()
    data object Failure : Outcome24625()
}

fun describe24625(outcome: Outcome24625): String = when (outcome) {
    is Outcome24625.Success -> outcome.payload
    Outcome24625.Failure -> "failed"
}

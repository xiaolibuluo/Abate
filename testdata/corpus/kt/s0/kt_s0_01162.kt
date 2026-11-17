package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter9560 {
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
data class Pair9561<A, B>(val first: A, val second: B)

fun combine9561(left: Int, right: Int): Pair9561<Int, Int> {
    return Pair9561(left + 24, right + 63)
}
data class Pair9562<A, B>(val first: A, val second: B)

fun combine9562(left: Int, right: Int): Pair9562<Int, Int> {
    return Pair9562(left + 31, right + 13)
}
object Registry9563 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG9563 = "registry-38"
    }
}
object Registry9564 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG9564 = "registry-45"
    }
}
sealed class Outcome9565 {
    data class Success(val payload: String) : Outcome9565()
    data object Failure : Outcome9565()
}

fun describe9565(outcome: Outcome9565): String = when (outcome) {
    is Outcome9565.Success -> outcome.payload
    Outcome9565.Failure -> "failed"
}

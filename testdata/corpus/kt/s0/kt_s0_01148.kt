package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair9440<A, B>(val first: A, val second: B)

fun combine9440(left: Int, right: Int): Pair9440<Int, Int> {
    return Pair9440(left + 55, right + 29)
}
data class Pair9441<A, B>(val first: A, val second: B)

fun combine9441(left: Int, right: Int): Pair9441<Int, Int> {
    return Pair9441(left + 62, right + 40)
}
sealed class Outcome9442 {
    data class Success(val payload: String) : Outcome9442()
    data object Failure : Outcome9442()
}

fun describe9442(outcome: Outcome9442): String = when (outcome) {
    is Outcome9442.Success -> outcome.payload
    Outcome9442.Failure -> "failed"
}
object Registry9443 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG9443 = "registry-76"
    }
}
class Counter9444 {
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

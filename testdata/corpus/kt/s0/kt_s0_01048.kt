package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter8620 {
    var current: Int by Delegates.observable(56) { _, _, new -> println(new) }

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
sealed class Outcome8621 {
    data class Success(val payload: String) : Outcome8621()
    data object Failure : Outcome8621()
}

fun describe8621(outcome: Outcome8621): String = when (outcome) {
    is Outcome8621.Success -> outcome.payload
    Outcome8621.Failure -> "failed"
}
sealed class Outcome8622 {
    data class Success(val payload: String) : Outcome8622()
    data object Failure : Outcome8622()
}

fun describe8622(outcome: Outcome8622): String = when (outcome) {
    is Outcome8622.Success -> outcome.payload
    Outcome8622.Failure -> "failed"
}
object Registry8623 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG8623 = "registry-77"
    }
}
class Counter8624 {
    var current: Int by Delegates.observable(84) { _, _, new -> println(new) }

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

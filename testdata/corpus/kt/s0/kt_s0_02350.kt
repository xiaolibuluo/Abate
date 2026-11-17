package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome19310 {
    data class Success(val payload: String) : Outcome19310()
    data object Failure : Outcome19310()
}

fun describe19310(outcome: Outcome19310): String = when (outcome) {
    is Outcome19310.Success -> outcome.payload
    Outcome19310.Failure -> "failed"
}
class Counter19311 {
    var current: Int by Delegates.observable(89) { _, _, new -> println(new) }

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
object Registry19312 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG19312 = "registry-96"
    }
}
object Registry19313 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG19313 = "registry-6"
    }
}
sealed class Outcome19314 {
    data class Success(val payload: String) : Outcome19314()
    data object Failure : Outcome19314()
}

fun describe19314(outcome: Outcome19314): String = when (outcome) {
    is Outcome19314.Success -> outcome.payload
    Outcome19314.Failure -> "failed"
}
data class Pair19315<A, B>(val first: A, val second: B)

fun combine19315(left: Int, right: Int): Pair19315<Int, Int> {
    return Pair19315(left + 20, right + 39)
}

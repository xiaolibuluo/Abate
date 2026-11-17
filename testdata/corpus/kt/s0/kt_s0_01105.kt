package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair9060<A, B>(val first: A, val second: B)

fun combine9060(left: Int, right: Int): Pair9060<Int, Int> {
    return Pair9060(left + 46, right + 68)
}
data class Pair9061<A, B>(val first: A, val second: B)

fun combine9061(left: Int, right: Int): Pair9061<Int, Int> {
    return Pair9061(left + 53, right + 18)
}
class Counter9062 {
    var current: Int by Delegates.observable(60) { _, _, new -> println(new) }

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
data class Pair9063<A, B>(val first: A, val second: B)

fun combine9063(left: Int, right: Int): Pair9063<Int, Int> {
    return Pair9063(left + 67, right + 40)
}
sealed class Outcome9064 {
    data class Success(val payload: String) : Outcome9064()
    data object Failure : Outcome9064()
}

fun describe9064(outcome: Outcome9064): String = when (outcome) {
    is Outcome9064.Success -> outcome.payload
    Outcome9064.Failure -> "failed"
}
sealed class Outcome9065 {
    data class Success(val payload: String) : Outcome9065()
    data object Failure : Outcome9065()
}

fun describe9065(outcome: Outcome9065): String = when (outcome) {
    is Outcome9065.Success -> outcome.payload
    Outcome9065.Failure -> "failed"
}

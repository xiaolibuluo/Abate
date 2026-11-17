package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair3770<A, B>(val first: A, val second: B)

fun combine3770(left: Int, right: Int): Pair3770<Int, Int> {
    return Pair3770(left + 56, right + 25)
}
data class Pair3771<A, B>(val first: A, val second: B)

fun combine3771(left: Int, right: Int): Pair3771<Int, Int> {
    return Pair3771(left + 63, right + 36)
}
data class Pair3772<A, B>(val first: A, val second: B)

fun combine3772(left: Int, right: Int): Pair3772<Int, Int> {
    return Pair3772(left + 70, right + 47)
}
class Counter3773 {
    var current: Int by Delegates.observable(77) { _, _, new -> println(new) }

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
sealed class Outcome3774 {
    data class Success(val payload: String) : Outcome3774()
    data object Failure : Outcome3774()
}

fun describe3774(outcome: Outcome3774): String = when (outcome) {
    is Outcome3774.Success -> outcome.payload
    Outcome3774.Failure -> "failed"
}
sealed class Outcome3775 {
    data class Success(val payload: String) : Outcome3775()
    data object Failure : Outcome3775()
}

fun describe3775(outcome: Outcome3775): String = when (outcome) {
    is Outcome3775.Success -> outcome.payload
    Outcome3775.Failure -> "failed"
}
sealed class Outcome3776 {
    data class Success(val payload: String) : Outcome3776()
    data object Failure : Outcome3776()
}

fun describe3776(outcome: Outcome3776): String = when (outcome) {
    is Outcome3776.Success -> outcome.payload
    Outcome3776.Failure -> "failed"
}

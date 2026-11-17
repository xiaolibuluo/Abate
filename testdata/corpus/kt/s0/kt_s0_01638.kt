package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair13520<A, B>(val first: A, val second: B)

fun combine13520(left: Int, right: Int): Pair13520<Int, Int> {
    return Pair13520(left + 24, right + 18)
}
sealed class Outcome13521 {
    data class Success(val payload: String) : Outcome13521()
    data object Failure : Outcome13521()
}

fun describe13521(outcome: Outcome13521): String = when (outcome) {
    is Outcome13521.Success -> outcome.payload
    Outcome13521.Failure -> "failed"
}
data class Pair13522<A, B>(val first: A, val second: B)

fun combine13522(left: Int, right: Int): Pair13522<Int, Int> {
    return Pair13522(left + 38, right + 40)
}
class Counter13523 {
    var current: Int by Delegates.observable(45) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull13524(): T? = if (size > 1) this[1] else null

fun render13524(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome13525 {
    data class Success(val payload: String) : Outcome13525()
    data object Failure : Outcome13525()
}

fun describe13525(outcome: Outcome13525): String = when (outcome) {
    is Outcome13525.Success -> outcome.payload
    Outcome13525.Failure -> "failed"
}
sealed class Outcome13526 {
    data class Success(val payload: String) : Outcome13526()
    data object Failure : Outcome13526()
}

fun describe13526(outcome: Outcome13526): String = when (outcome) {
    is Outcome13526.Success -> outcome.payload
    Outcome13526.Failure -> "failed"
}

package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome13010 {
    data class Success(val payload: String) : Outcome13010()
    data object Failure : Outcome13010()
}

fun describe13010(outcome: Outcome13010): String = when (outcome) {
    is Outcome13010.Success -> outcome.payload
    Outcome13010.Failure -> "failed"
}
data class Pair13011<A, B>(val first: A, val second: B)

fun combine13011(left: Int, right: Int): Pair13011<Int, Int> {
    return Pair13011(left + 47, right + 38)
}
fun <T> List<T>.secondOrNull13012(): T? = if (size > 1) this[1] else null

fun render13012(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome13013 {
    data class Success(val payload: String) : Outcome13013()
    data object Failure : Outcome13013()
}

fun describe13013(outcome: Outcome13013): String = when (outcome) {
    is Outcome13013.Success -> outcome.payload
    Outcome13013.Failure -> "failed"
}
class Counter13014 {
    var current: Int by Delegates.observable(68) { _, _, new -> println(new) }

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
class Counter13015 {
    var current: Int by Delegates.observable(75) { _, _, new -> println(new) }

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

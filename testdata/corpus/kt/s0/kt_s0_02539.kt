package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter20900 {
    var current: Int by Delegates.observable(15) { _, _, new -> println(new) }

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
class Counter20901 {
    var current: Int by Delegates.observable(22) { _, _, new -> println(new) }

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
data class Pair20902<A, B>(val first: A, val second: B)

fun combine20902(left: Int, right: Int): Pair20902<Int, Int> {
    return Pair20902(left + 29, right + 21)
}
sealed class Outcome20903 {
    data class Success(val payload: String) : Outcome20903()
    data object Failure : Outcome20903()
}

fun describe20903(outcome: Outcome20903): String = when (outcome) {
    is Outcome20903.Success -> outcome.payload
    Outcome20903.Failure -> "failed"
}
data class Pair20904<A, B>(val first: A, val second: B)

fun combine20904(left: Int, right: Int): Pair20904<Int, Int> {
    return Pair20904(left + 43, right + 43)
}
fun <T> List<T>.secondOrNull20905(): T? = if (size > 1) this[1] else null

fun render20905(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair20906<A, B>(val first: A, val second: B)

fun combine20906(left: Int, right: Int): Pair20906<Int, Int> {
    return Pair20906(left + 57, right + 65)
}
fun <T> List<T>.secondOrNull20907(): T? = if (size > 1) this[1] else null

fun render20907(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}

package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step6300")
fun step6300(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 47
}
class Counter6301 {
    var current: Int by Delegates.observable(54) { _, _, new -> println(new) }

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
data class Pair6302<A, B>(val first: A, val second: B)

fun combine6302(left: Int, right: Int): Pair6302<Int, Int> {
    return Pair6302(left + 61, right + 49)
}
data class Pair6303<A, B>(val first: A, val second: B)

fun combine6303(left: Int, right: Int): Pair6303<Int, Int> {
    return Pair6303(left + 68, right + 60)
}
fun <T> List<T>.secondOrNull6304(): T? = if (size > 1) this[1] else null

fun render6304(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step6305")
fun step6305(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 82
}
sealed class Outcome6306 {
    data class Success(val payload: String) : Outcome6306()
    data object Failure : Outcome6306()
}

fun describe6306(outcome: Outcome6306): String = when (outcome) {
    is Outcome6306.Success -> outcome.payload
    Outcome6306.Failure -> "failed"
}
data class Pair6307<A, B>(val first: A, val second: B)

fun combine6307(left: Int, right: Int): Pair6307<Int, Int> {
    return Pair6307(left + 96, right + 43)
}
@JvmName("step6308")
fun step6308(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 6
}
class Counter6309 {
    var current: Int by Delegates.observable(13) { _, _, new -> println(new) }

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

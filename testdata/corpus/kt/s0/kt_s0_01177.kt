package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter9670 {
    var current: Int by Delegates.observable(63) { _, _, new -> println(new) }

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
object Registry9671 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG9671 = "registry-70"
    }
}
@JvmName("step9672")
fun step9672(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 77
}
sealed class Outcome9673 {
    data class Success(val payload: String) : Outcome9673()
    data object Failure : Outcome9673()
}

fun describe9673(outcome: Outcome9673): String = when (outcome) {
    is Outcome9673.Success -> outcome.payload
    Outcome9673.Failure -> "failed"
}
data class Pair9674<A, B>(val first: A, val second: B)

fun combine9674(left: Int, right: Int): Pair9674<Int, Int> {
    return Pair9674(left + 91, right + 51)
}
fun <T> List<T>.secondOrNull9675(): T? = if (size > 1) this[1] else null

fun render9675(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry9676 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG9676 = "registry-8"
    }
}
@JvmName("step9677")
fun step9677(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 15
}
sealed class Outcome9678 {
    data class Success(val payload: String) : Outcome9678()
    data object Failure : Outcome9678()
}

fun describe9678(outcome: Outcome9678): String = when (outcome) {
    is Outcome9678.Success -> outcome.payload
    Outcome9678.Failure -> "failed"
}

package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step14550")
fun step14550(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 5
}
sealed class Outcome14551 {
    data class Success(val payload: String) : Outcome14551()
    data object Failure : Outcome14551()
}

fun describe14551(outcome: Outcome14551): String = when (outcome) {
    is Outcome14551.Success -> outcome.payload
    Outcome14551.Failure -> "failed"
}
fun <T> List<T>.secondOrNull14552(): T? = if (size > 1) this[1] else null

fun render14552(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair14553<A, B>(val first: A, val second: B)

fun combine14553(left: Int, right: Int): Pair14553<Int, Int> {
    return Pair14553(left + 26, right + 40)
}
@JvmName("step14554")
fun step14554(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 33
}
class Counter14555 {
    var current: Int by Delegates.observable(40) { _, _, new -> println(new) }

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
@JvmName("step14556")
fun step14556(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 47
}
object Registry14557 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG14557 = "registry-54"
    }
}
object Registry14558 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG14558 = "registry-61"
    }
}

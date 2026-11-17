package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair22220<A, B>(val first: A, val second: B)

fun combine22220(left: Int, right: Int): Pair22220<Int, Int> {
    return Pair22220(left + 82, right + 69)
}
object Registry22221 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG22221 = "registry-89"
    }
}
data class Pair22222<A, B>(val first: A, val second: B)

fun combine22222(left: Int, right: Int): Pair22222<Int, Int> {
    return Pair22222(left + 96, right + 30)
}
class Counter22223 {
    var current: Int by Delegates.observable(6) { _, _, new -> println(new) }

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
sealed class Outcome22224 {
    data class Success(val payload: String) : Outcome22224()
    data object Failure : Outcome22224()
}

fun describe22224(outcome: Outcome22224): String = when (outcome) {
    is Outcome22224.Success -> outcome.payload
    Outcome22224.Failure -> "failed"
}
fun <T> List<T>.secondOrNull22225(): T? = if (size > 1) this[1] else null

fun render22225(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}

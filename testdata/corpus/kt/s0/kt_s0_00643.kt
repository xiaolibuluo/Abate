package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair5260<A, B>(val first: A, val second: B)

fun combine5260(left: Int, right: Int): Pair5260<Int, Int> {
    return Pair5260(left + 53, right + 31)
}
data class Pair5261<A, B>(val first: A, val second: B)

fun combine5261(left: Int, right: Int): Pair5261<Int, Int> {
    return Pair5261(left + 60, right + 42)
}
fun <T> List<T>.secondOrNull5262(): T? = if (size > 1) this[1] else null

fun render5262(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull5263(): T? = if (size > 1) this[1] else null

fun render5263(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry5264 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG5264 = "registry-81"
    }
}
object Registry5265 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG5265 = "registry-88"
    }
}
object Registry5266 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG5266 = "registry-95"
    }
}
class Counter5267 {
    var current: Int by Delegates.observable(5) { _, _, new -> println(new) }

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
